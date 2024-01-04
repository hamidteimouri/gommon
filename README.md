# Gommon (Golang Common)

Gommon is a golang common package that provides several methods to use in your projects.<br>
Gommon has several sub-package :

- htblockchain
- htcolog ( colorful console logger )
- htenvier ( ENV methods )
- htregex  ( work with regex )

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

### Format
Format email , username and ...

example :

```go
gommon.MakeMaskEmailAndDomain("usernaame@gmail.com")    // output : use***@gm**.com
gommon.MakeMaskEmail("usernaame@gmail.com")             // output : use***@gmail.com
gommon.MakeMaskUsername("myusername")                   // output : my***me
```

### htregex

example :

```go
htregex.IsEmail("username@gmail.com")
```
