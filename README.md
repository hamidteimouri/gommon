# Gommon (Golang Common)

Gommon is a golang common package that provides several methods to use in your projects.<br>
Gommon has several sub-package :

- htapplife
- htblockchain
- htcolog ( colorful console logger )
- htenvier ( ENV methods )
- htformat ( change format of a string like email , ... )
- htrandom ( random generator )
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

### htformat

example :

```go
htformat.MakeEmailInvisible("exampleEmail@test.com") // output : exa***@test.com
htformat.MakeUsernameInvisible("myusername") // output : my***me
```

### htregex

example :

```go
htregex.IsEmail("username@gmail.com")
```
