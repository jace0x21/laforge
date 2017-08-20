###########################################################
# LAFORGE GENERATED TERRAFORM CONFIGURATION
# -- WARNING --
# This file is automatically generated. Do not edit it's
# contents directly. Use Laforge and re build this file.
###########################################################

# AWS Configuration
provider "aws" {
  access_key = "{{ .Competition.AWSCred.APIKey }}"
  secret_key = "{{ .Competition.AWSCred.APISecret }}"
  region = "{{ .Environment.AWSConfig.Region }}"
  profile = ""
}

# Ubuntu AMI
data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-trusty-16.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

# Key Pair
resource "aws_key_pair" "ssh_{{ .Environment.Name }}" {
  key_name = "ssh_{{ .Environment.Name }}"
  public_key = "{{ .Competition.SSHPublicKey | html }}"
}

{{range $i, $_ := N .Environment.PodCount }}

{{ $id := printf "%s%d" $.Environment.Prefix $i }}

###########################################################
# BEGIN TEAM BLOCK
# team:{{ $id }}
###########################################################

# vpc:{{ $id }}
resource "aws_vpc" "{{ $id }}_vpc" {
  cidr_block = "{{ $.Environment.AWSConfig.CIDR }}"
  enable_dns_hostnames = true

  tags {
    Name = "{{ $id }}_vpc"
    Environment = "{{ $.Environment.Name }}"
    Team = "{{ $id }}"
  }
}

# igw:{{ $id }}
resource "aws_internet_gateway" "{{ $id }}_igw" {
  vpc_id = "${aws_vpc.{{ $id }}_vpc.id}"
}

# route_table:{{ $id }}
resource "aws_route_table" "{{ $id }}_default_gateway" {
  vpc_id = "${aws_vpc.{{ $id }}_vpc.id}"

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway.{{ $id }}_igw.id}"
  }

  tags {
    Name = "{{ $id }}_igw_route"
    Environment = "{{ $.Environment.Name }}"
    Team = "{{ $id }}"
  }
}

# vdi_subnet:{{ $id }}
resource "aws_subnet" "{{ $id }}_vdi_subnet" {
  vpc_id     = "${aws_vpc.{{ $id }}_vpc.id}"
  cidr_block = "{{ $.Environment.JumpHosts.CIDR }}"

  tags {
    Name = "{{ $id }}_vdi_subnet"
    Environment = "{{ $.Environment.Name }}"
    Team = "{{ $id }}"
  }
}

# vdi_route:{{ $id }}
resource "aws_route_table_association" "{{ $id }}_rt_assoc" {
  subnet_id      = "${aws_subnet.{{ $id }}_vdi_subnet.id}"
  route_table_id = "${aws_route_table.{{ $id }}_default_gateway.id}"
}

# r53:{{ $id }}
resource "aws_route53_zone" "{{ $id }}_r53" {
  name = "{{ $.Environment.Domain }}"
  description = "Laforge - {{ $id }}"
  vpc_id = "${aws_vpc.{{ $id }}_vpc.id}"

  tags {
    Environment = "{{ $.Environment.Name }}"
    Team = "{{ $id }}"
  }
}

# r53_assoc:{{ $id }}
resource "aws_route53_zone_association" "{{ $id }}_r53_assoc" {
  zone_id = "${aws_route53_zone.{{ $id }}_r53.zone_id}"
  vpc_id  = "${aws_vpc.{{ $id }}_vpc.id}"
}

# sg_admin:{{ $id }}
resource "aws_security_group" "{{ $id }}_admin" {
  name = "{{ $id }}_admin"
  description = "Laforge - {{ $id }}_admin SG"
  vpc_id = "${aws_vpc.{{ $id }}_vpc.id}"

  tags {
    Name = "{{ $id }}_admin"
    Environment = "{{ $.Environment.Name }}"
    Team = "{{ $id }}"
  }
}

# sg_base:{{ $id }}
resource "aws_security_group" "{{ $id }}_base" {
  name = "{{ $id }}_base"
  description = "Laforge - {{ $id }}_base SG"
  vpc_id = "${aws_vpc.{{ $id }}_vpc.id}"

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = [
      {{range $_, $AdminIP := $.Competition.AdminIPs }}
      "{{ $AdminIP }}",
      {{end}}
      # FIX PUBLIC IP: "{{/* MyIP */}}/32",
    ]
  }

  ingress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    security_groups = [
      "${aws_security_group.{{ $id }}_admin.id}",
    ]
  }

  ingress {
    from_port = 8
    to_port = 0
    protocol = "icmp"
    self = true
  }

  tags {
    Name = "{{ $id }}_base"
    Environment = "{{ $.Environment.Name }}"
    Team = "{{ $id }}"
  }
}

# sg_vdi:{{ $id }}
resource "aws_security_group" "{{ $id }}_vdi" {
  name = "{{ $id }}_vdi"
  description = "Laforge - {{ $id }}_vdi SG"
  vpc_id = "${aws_vpc.{{ $id }}_vpc.id}"

  ingress {
    from_port         = 22
    to_port           = 22
    protocol          = "tcp"
    cidr_blocks       = [
      {{range $_, $whitelistIP := $.Environment.WhitelistIPs }}
      "{{ $whitelistIP }}",
      {{end}}
    ]
  }

  ingress {
    from_port         = 3389
    to_port           = 3389
    protocol          = "tcp"
    cidr_blocks       = [
      {{range $_, $whitelistIP := $.Environment.WhitelistIPs }}
      "{{ $whitelistIP }}",
      {{end}}
    ]
  }

  ingress {
    from_port         = 5900
    to_port           = 5900
    protocol          = "tcp"
    cidr_blocks       = [
      {{range $_, $whitelistIP := $.Environment.WhitelistIPs }}
      "{{ $whitelistIP }}",
      {{end}}
    ]
  }

  ingress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    security_groups = [
      "${aws_security_group.{{ $id }}_base.id}",
    ]
  }

  tags {
    Name = "{{ $id }}_vdi"
    Environment = "{{ $.Environment.Name }}"
    Team = "{{ $id }}"
  }
}


# sg_rule_admin:{{ $id }}
resource "aws_security_group_rule" "{{ $id }}_admin_callback" {
  type              = "ingress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  source_security_group_id = "${aws_security_group.{{ $id }}_base.id}"
  security_group_id = "${aws_security_group.{{ $id }}_admin.id}"
}

{{/* TODO: VDI Network Creation */}}

{{range $wjh, $_ := N $.Environment.JumpHosts.Windows.Count }}

{{ $hostname := printf "%s-jump-win-%d" $id $wjh }}

resource "aws_instance" "{{ $hostname }}" {

  ami = "{{ $.Environment.WindowsJumpAMI }}"
  instance_type = "{{ $.Environment.JumpHosts.Windows.Size }}"
  key_name = "${aws_key_pair.ssh_{{ $.Environment.Name }}}"
  subnet_id = "${aws_subnet.{{ $id }}_vdi_subnet.id}"

  {{ $customIP := CustomIP $.Environment.JumpHosts.CIDR 20 $wjh }}

  private_ip = "{{ $customIP }}"

  vpc_security_group_ids = [
    "${aws_security_group.{{ $id }}_base.id}",
    "${aws_security_group.{{ $id }}_{{ $id }}_vdi.id}",
  ]

  tags {
    Name = "{{ $hostname }}"
    Environment = "{{ $.Environment.Name }}"
    Network = "{{ $id }}-vdi"
    Team = "{{ $id }}"
  }
}

{{end}}

{{range $wjh, $_ := N $.Environment.JumpHosts.Kali.Count }}

{{ $hostname := printf "%s-jump-kali-%d" $id $wjh }}

resource "aws_instance" "{{ $hostname }}" {

  ami = "{{ $.Environment.KaliJumpAMI }}"
  instance_type = "{{ $.Environment.JumpHosts.Kali.Size }}"
  key_name = "${aws_key_pair.ssh_{{ $.Environment.Name }}}"
  subnet_id = "${aws_subnet.{{ $id }}_vdi_subnet.id}"

  {{ $customIP := CustomIP $.Environment.JumpHosts.CIDR 30 $wjh }}

  private_ip = "{{ $customIP }}"

  vpc_security_group_ids = [
    "${aws_security_group.{{ $id }}_base.id}",
    "${aws_security_group.{{ $id }}_{{ $id }}_vdi.id}",
  ]

  tags {
    Name = "{{ $hostname }}"
    Environment = "{{ $.Environment.Name }}"
    Network = "{{ $id }}-vdi"
    Team = "{{ $id }}"
  }
}

{{end}}


{{range $port, $_ := $.Environment.ResolvePublicTCP }}

resource "aws_security_group" "{{ $id }}_public_sg_tcp_{{ $port }}" {
  name = "{{ $id }}_public_sg_tcp_{{ $port }}"
  description = "Laforge - {{ $id }}_public tcp/{{ $port }} SG"
  vpc_id = "${aws_vpc.{{ $id }}_vpc.id}"

  ingress {
    from_port = {{ $port }}
    to_port = {{ $port }}
    protocol = "tcp"
    security_groups = [
      "${aws_security_group.{{ $id }}_base.id}",
    ]
  }

  tags {
    Name = "{{ $id }}_public_sg_tcp_{{ $port }}"
    Environment = "{{ $.Environment.Name }}"
    Team = "{{ $id }}"
  }
}

{{end}}

{{range $port, $_ := $.Environment.ResolvePublicUDP }}

resource "aws_security_group" "{{ $id }}_public_sg_udp_{{ $port }}" {
  name = "{{ $id }}_public_sg_udp_{{ $port }}"
  description = "Laforge - {{ $id }}_public udp/{{ $port }} SG"
  vpc_id = "${aws_vpc.{{ $id }}_vpc.id}"

  ingress {
    from_port = {{ $port }}
    to_port = {{ $port }}
    protocol = "udp"
    security_groups = [
      "${aws_security_group.{{ $id }}_base.id}",
    ]
  }

  tags {
    Name = "{{ $id }}_public_sg_tcp_{{ $port }}"
    Environment = "{{ $.Environment.Name }}"
    Team = "{{ $id }}"
  }
}

{{end}}


{{range $_, $network := $.Environment.ResolvedNetworks }}

# subnet:{{ $network.Name }}:{{ $id }}
resource "aws_subnet" "{{ $id }}_{{ $network.Name }}" {
  vpc_id     = "${aws_vpc.{{ $id }}_vpc.id}"
  cidr_block = "{{ $network.CIDR }}"

  tags {
    Name = "{{ $id }}_{{ $network.Name }}"
    Environment = "{{ $.Environment.Name }}"
    Team = "{{ $id }}"
  }
}


# rt_assoc:{{ $network.Name }}:{{ $id }}
resource "aws_route_table_association" "{{ $id }}_{{ $network.Name }}_rt_assoc" {
  subnet_id      = "${aws_subnet.{{ $id }}_{{ $network.Name }}.id}"
  route_table_id = "${aws_route_table.{{ $id }}_default_gateway.id}"
}

# sg:{{ $network.Name }}:{{ $id }}
resource "aws_security_group" "{{ $id }}_{{ $network.Name }}" {
  name = "{{ $id }}_{{ $network.Name }}"
  description = "Laforge - {{ $id }}_{{ $network.Name }} SG"
  vpc_id = "${aws_vpc.{{ $id }}_vpc.id}"

  ingress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    self = true
  }

  {{if $network.VDIVisible }}
  ingress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    security_groups = [
      "${aws_security_group.{{ $id }}_vdi.id}"
    ]
  }
  {{end}}

  tags {
    Name = "{{ $id }}_{{ $network.Name }}"
    Environment = "{{ $.Environment.Name }}"
    Team = "{{ $id }}"
  }
}

{{range $hostname, $host := $network.ResolvedHosts }}

{{ $hostIP := CustomIP $network.CIDR 0 $host.LastOctet }}

# instance:{{ $id }}_{{ $network.Subdomain }}_{{ $hostname }}
resource "aws_instance" "{{ $id }}_{{ $network.Subdomain }}_{{ $hostname }}" {

  ami = "${data.aws_ami.ubuntu.id}"
  instance_type = "{{ $host.InstanceSize }}"
  key_name = "${aws_key_pair.ssh_{{ $.Environment.Name }}}"
  subnet_id = "${aws_subnet.{{ $id }}_{{ $network.Name }}.id}"

  private_ip = "{{ $hostIP }}"

  vpc_security_group_ids = [
    "${aws_security_group.{{ $id }}_base.id}",
    "${aws_security_group.{{ $id }}_{{ $network.Name }}.id}",
    {{ range $_, $port := $host.TCPPorts }}
      "${aws_security_group.{{ $id }}_public_sg_tcp_{{ $port }}.id}",
    {{end}}
    {{ range $_, $port := $host.UDPPorts }}
      "${aws_security_group.{{ $id }}_public_sg_udp_{{ $port }}.id}",
    {{end}}
  ]

  tags {
    Name = "{{ $id }}_{{ $network.Subdomain }}_{{ $hostname }}"
    Hostname = "{{ $hostname }}"
    Environment = "{{ $.Environment.Name }}"
    Network = "{{ $network.Name }}"
    Team = "{{ $id }}"
  }

  {{if eq $host.OS "windows" }}
    {{ $script_path := ScriptRender "windows_uds.xml" $.Competition $.Environment $i $network $host $hostname}}


    user_data = "${file("{{ $script_path }}")}"

    connection {
      type     = "winrm"
      user     = "Administrator"
      password = "{{ $.Competition.RootPassword }}"
    }
  {{end}}

  {{if eq $host.OS "ubuntu" }}
    {{ $script_path := ScriptRender "ubuntu_uds.sh" $.Competition $.Environment $i $network $host $hostname}}


    user_data = "${file("{{ $script_path }}")}"

    connection {
      type     = "ssh"
      user     = "root"
      private_key = "${file("{{ $.Competition.SSHPrivateKeyPath }}")}"
    }

  {{end}}

  {{if eq $host.OS "centos" }}

    connection {
      type     = "ssh"
      user     = "root"
      private_key = "${file("{{ $.Competition.SSHPrivateKeyPath }}")}"
    }

  {{end}}

  {{if eq $host.OS "freebsd" }}

    connection {
      type     = "ssh"
      user     = "root"
      private_key = "${file("{{ $.Competition.SSHPrivateKeyPath }}")}"
    }

  {{end}}


  provisioner "remote-exec" {
    scripts = [
      {{range $_, $sname := $host.Scripts }}
        {{ $script_path := DScript $sname $.Competition $.Environment $i $network $host $hostname }}
        "${file("{{ $script_path }}"}",
      {{end}}
    ]
  }
}

{{ range $_, $cname := $host.InternalCNAMEs }}

{{ $recordValue := CustomInternalCNAME $.Environment $network $cname }}

resource "aws_route53_record" "{{ $id }}_icname_{{ $cname }}" {
  zone_id = "${aws_route53_zone.{{ $id }}_r53.zone_id}"
  name    = "{{ $recordValue }}"
  type    = "A"
  ttl     = "60"
  records = [
    "{{ $hostIP }}"
  ]
}

{{end}} {{/* End ExternalCNAME Iterator */}}

{{ range $_, $cname := $host.ExternalCNAMEs }}

{{ $recordValue := CustomExternalCNAME $.Environment $cname }}

resource "aws_route53_record" "{{ $id }}_ecname_{{ $cname }}" {
  zone_id = "${aws_route53_zone.{{ $id }}_r53.zone_id}"
  name    = "{{ $recordValue }}"
  type    = "A"
  ttl     = "60"
  records = [
    "{{ $hostIP }}"
  ]
}

{{end}} {{/* End InternalCNAME Iterator */}}

{{end}} {{/* End Host Iterator */}}
{{end}} {{/* End Network Iterator */}}
{{end}} {{/* End Pod Iterator */}}