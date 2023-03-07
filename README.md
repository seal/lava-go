# lava-go
Simple program that generates lavalamp-blobs in terminal 


### Stuff to add:

- Optional flags for speed, radius, etc, more customisation 
- Changing colors / using ascii signs 


## Description

Program inspired by https://github.com/AngelJumbo/lavat 

CLI program with TUI that makes lavalamp blob like objects 

![lavagofinal](https://user-images.githubusercontent.com/25641834/205149258-3bebe79a-1bfe-494f-9794-4511dd4557cd.gif)



## How To Use
```
go run main.go 
```

##Flags:

-b [int] Number of balls

-r [int] Radius of balls

-s [int] Polling rate in microseconds, default 50000

-c [0/1] If the balls are contained in a box

Exit the program by pressing Esc and waiting up to 5 seconds 



## Credit

- Original idea + math - https://github.com/AngelJumbo/lavat  
- Termbox - https://github.com/nsf/termbox-go 



