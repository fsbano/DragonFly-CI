# DragonFly-CI
DSL - NextGeneration

```
With '/etc/ssh/sshd_config'
  do
    Set-Variable = (
      X11Forwarding => no,
      IgnoreRhosts => yes,
      MaxAuthTries => 4,
      HostbasedAuthentication => no,
      PermitRootLogin => no,
      PermitEmptyPasswords => no,
      PermitUserEnvironment => no,
      ClientAliveInterval => 300,
      ClientAliveCountMax => 0,
      DenyUsers => nobody,
      DenyGroups => nobody,
      AllowUsers => guest
    )
    Absent-Variable = (
      PrintMotd,
      PrintLastLog,
      TCPKeepAlive
    )
    Comment-Variable = (
      UseDNS
    )
    Set-Permission = 0700
    Restart => True
done
```
