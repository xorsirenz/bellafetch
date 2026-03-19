package linux

func GetLinuxData() map[string]any {

		var data = make(map[string]any)
		
		data["user"] = Username()
		data["host"] = Hostname()
		data["prettyname"] = PrettyName()
		data["kernel"] = Kernel()
		data["uptime"] = Uptime()
		data["pkgs"] = PkgManager()
		data["wm"] = ""
		data["cpu"] = Cpu()
		data["gpu"] = Gpu()
		data["diskSpace"] = Storage()
		data["memory"] = Memory()
		return data
}
