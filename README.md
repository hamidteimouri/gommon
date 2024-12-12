# Gommon (Golang Common)

`gommon` is a golang common package that provides several methods to use in your projects.<br>
`gommon` has several sub-package :


- htcolog ( colorful console logger )
- htenvier ( ENV methods )

### htcolog

example :

```go
htcolog.DoGreen("prints a green string")
htcolog.MakeRed("returns a red string")
```

### htenvier

example :

```go
htenvier.ENV("DB_NAME")
```

### Format & Validation
Format email , username and ...

example :

```go
gommon.MakeMaskEmailAndDomain("usernaame@gmail.com")    // output : use***@gm**.com
gommon.MakeMaskEmail("usernaame@gmail.com")             // output : use***@gmail.com
gommon.MakeMaskUsername("myusername")                   // output : my***me
gommon.IsEmail("username@gmail.com")
```
