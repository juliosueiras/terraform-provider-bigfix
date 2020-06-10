variable "fixlet_name"{
  description = "Name of fixlet"
}

variable "linux-vm-name"{
  description = "Name of computer to patch"
}

variable "linux-sites"{
  description = "List of sites from which you want to patch"
}

variable "bigfix_username"{
    description = "BES console operator username with appropriate access rights"
}

variable "bigfix_password"{
    description = "Password of BES user"
}

variable "bigfix_server"{
    description = "IP address or name of bigfix server"
}

variable "bigfix_port"{
    description = "Port number for BES rest service"
    default=52311
}