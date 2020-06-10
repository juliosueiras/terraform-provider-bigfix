# Configure the Bigfix Provider
provider "bigfix" {
  port = var.bigfix_port
  username = var.bigfix_username
  password = var.bigfix_password
  server = var.bigfix_server
}

# data source bigfix fixlet
data "bigfix_fixlet" "myfixlet"{
	name = var.fixlet_name
}

# data source bigfix computer
data  "bigfix_computer" "linux_vm"{
 	name = var.linux-vm-name
}

# resource bigfix multiple_action_group
resource "bigfix_multiple_action_group" "test" {
  input_file_name = "MAG.xml" 
  target_computer_id = data.bigfix_computer.linux_vm.id
  site_name = var.linux-sites
}	

output "MAG-linux-vm" {
  value = bigfix_multiple_action_group.test.title
}

output "action-state-linux" {
  value = bigfix_multiple_action_group.test.state
}

