package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
	"strings"
	)

func process(direction int, text string) {
	txt_arr := []rune(text)
	txt_arr = txt_arr[:len(txt_arr)-1]
	str_len := len(txt_arr)
	var coeff int64 = 0
	delay := time.Duration(coeff * int64(time.Millisecond))
	
	switch direction {
	
	case 1:
		for i := 0; i < str_len; i++ {
			fmt.Printf("%c", txt_arr[i])
			time.Sleep(delay)
		}
		fmt.Printf("\n")
	
	case 2:
		for i := str_len-1; i >= 0; i-- {
			if txt_arr[i] == '\r' {
				fmt.Print("kuku!")
			}
			fmt.Printf("%c", txt_arr[i])
			time.Sleep(delay)
		}
		fmt.Printf("\n")
	
	
	}
}

func image(name string) {	
	
	name = strings.ReplaceAll(name, "\n", "")
	name = strings.ReplaceAll(name, "\r", "")
	for len(name) < 59 {
    name += " "
}
	name = "           @@@@@@@   @@@@@@@@@@@@@@@@@@@{               Hello " + name + "\n" 
	
	var img = "" +
"                                     ]<                                                                                   \n"+
"                                    @@@                                                                                   \n"+
"                                  =@@@@.                                                                                  \n"+
"                                 ]@@@@@.                                                                                  \n"+
"                                )@@@@@@                                                                                   \n"+
"                               =@@@@@@^                                                                                   \n"+
"                               @@@@@@@                                                                                    \n"+
"                              ]@@@@@@<                                                                                   \n"+
"                              {@@@@@@                                                                                     \n"+
"                              @@@@@@                                                                                      \n"+
"                             =@@@@@                     )@@}                                                              \n"+
"                             ^@@@@.                +%@@@@]                                                                \n"+
"                             ^@@@@            +%@@@@@@@                                                                   \n"+
"                             ^@@@           @@@@@@@@=                                                                      \n"+
"                             {@@         ]@@@@@}-                                                                          \n"+
"                   <@@@={  -{@@@      +@@@@+                          -.                                   -.              \n"+
"                 =@@@@@@@@@@@@@@@@=-@@@)             {@@<   @@%   =@@@)  {@@    @@%        <@@<        =@%   %@)          \n"+
"                 ^@@@@@@@@@@@@@@@@@@@@+              .@@    @@=    @@%  .@@@    @@=         @@        ^@@     @@@         \n"+
"                 @@@@@@@@@@@@@@@@@@@@@@@             .@@   +@@=    @@{ ^        @@=         @@        @@)     )@@         \n"+
"                 <@@@@@@@@@@@@@@@@@@@@)^{            -@@%@@.@@=    @@@%@        @@=         @@        @@)  )  +@@=        \n"+
"                   @@@@@@@@@@@@@@@@@@@^-             -@@    @@=    @@{          @@=    @    @@    @@  @@]     }@@         \n"+
"                 +} @@@@@@@@@@@@@  @@@               -@@    @@=    @@@    @)    @@@   ]@    @@<   @@   @@     @@-         \n"+
"                    {@@@@@@@@@@@@@@@@                @@@%  ^@@@   =@@@@@@@@@-  ^@@@@@@@@@  {@@@@@@@@    <@}-)@]           \n"+
"                     =@@@@@@@@@@@@@@@@                                                                                   \n"+
"                       ]@@@@@@@@@@@@%                                                                                     \n"+
"                       =@@@@@@@@@@@@                                                                                      \n"+
"                       @@@@@@@)-<{{                                                                                        \n"+
"                      .@@@@@@@                                                                                             \n"+
"                     )@@@@@@@{                    .  - =  }@^      )}]}+        <%@%)     =            .]@@{^       - -    \n"+
"                   {@@@@@@@@@@.                }@@   ]@@ @@@@]   @@    }@)   .@@@@+-]@@    @@}       -@+-+  ]@@     %@@   \n"+
"              <@@@@@@@@@@@@@@@@@@%]^            @@+  {@@    @^  @@]     @@<   @@@ ^< ]@@   @@<       @ @@%    @@    <@<   \n"+
"             @@@@@@@@@@@@@@@@@@@@@@@@@.         =@@  @@@}  {]  ^@@   -  <@@   @@@.} =@@    @@<         }@]    @@@   -@-   \n"+
"            @@@@@@@@@@@@@@@@@@@@@@@@@@@          @@=}^<@@ ])   <@@      <@@   @@{ }@@^     @@<   +=    <@]    @@@    @    \n"+
"            @@@@@@@@@@@@@@@@@@@@@)@@@@@=         =@@@  @@@}     @@)     @@)   @@{  +@@     @@<    @    <@]    @@<    %    \n"+
"            @@@@@@@@@@@@@@@@@@@@@@@@@@@@@^        @@}  )@@=      @@    )@]    @@%   <@@    @@@=  @@    ]@@])@@@}     ^    \n"+
"            @@@@@@@@@@@@@@@@@@@@@%@@@@@@@@@}      ]])  )]]+        }{]}^     .]])-  =]]]. =]]]]]]]}}   ]]]]]<.       @    \n"+
"            @@@@@@@@@@@@@@@@@@@@@@@@@{@@@@@@                                                                              \n"+
"            @@@@@@}@@@@@@@@@@@@@@@@@@%@@@@@                                                                              \n"+
"           {@@@@@@ -@@@@@@@@@@@@@@@@@@@@@<                                                                               \n"+
name +
"           @@@@@@=    @@@@@@@@@@@@@@@@@@@<                                                                               \n"+
"           @@@@@@      @@@@@@@@@@@@@@@@@@<       Glad that you are reading this!                                        \n"+
"           @@@@@]       %@@@@@@@@@@@@@@@{.       I have played a little with strings, switches and other stuff          \n"+
"          ]@@@@@.        }@@@@@@@@@@@@@@%        Some things work unusual, like, it is not so easy for the start       \n"+
"          @@@@@@          @@@@@@@@@@@@@@)        to detect int or float in a string with trailing chars,               \n"+
"         )@@@@@@           @@@@@@@@@@@@@^        and so on.                                                            \n"+
"         @@@@@@@          {@@@@@@@@@@@@@         Some \\r  and \\n chars are difficult to catch as well.                \n"+
"         @@@@@@<         %@@@@@@@@@@@@@@         So I was too lazy to do all checks, who cares )                      \n"+
"         @@@@@@        =@@@@@@@@@@@@@@@@         Anyway, Go is super cool to use, in my opinion.                      \n"+
"         @@@@@}      =@@@@@@@@@@@@@@@@@@+                                                                             \n"+
"         @@@@@      @@@@@@@@@@@@@@@@@@@@)        And thanks a lot for the course!                                     \n"+
"         @@@@{    ^@@@@@@@@@@@@@@@@@@@@@         It unites people, that is super good!                                \n"+
"         @@@@.   ]@@@@@@@@@@@@@@@@@@@@@@         See you around!                                                      \n"+
"         @@@@   =@@@@@@@@@@@@@@@@@@%@@@@                                                                              \n"+
"         @@@)   @@@@@@@@@@@@@@@@@@@@@@%@+                                                                             \n"+
"         @@@.   @@@@@@@@@@@@@@@@@@@@@{@@@                                                                             \n"+
"        =@@@    @@@@@@@@@@@@@@@@@@@@=]@@@@                                                                             \n"+
"        @@@@-   @@@@@@@@@@@@@@@@@@@@@@@@@@%                                                   @dohtur03                \n"+
"       %@@@@    )@@@@@@@@@@@@@@@@@@@@]@@@@@.                                                                            \n"+
"      @@@@@@     @@@@@@@@@@@@@@@@@@@@@@@@@@@                                                  school_21                 \n"+
"     <@% =@       @@@@@@@@@@@@@@@@@@@@@@@@@@+                                                                           \n"+
"    }@{  %@       <@@@@@@@@@@@@@@@@@@@@@@@@@@                                                                           \n"+
"  <@]+            +@@@@@@@@@@@@@@@@@@%^)^-                                                                           \n"+
"                      .=+^^^^=-.                                                                           \n"

	
	process(1, img)
	process(2, img)
}

func main() {
	
	buf_reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Hello world!    ₍^. .^₎\nPlease open full-screen and enter your name: ")
	name, err := buf_reader.ReadString('\n')
	
	if name == "" || err != nil {
		fmt.Print("Something went wrong")
	} else {
		image(name)
	} 
		
}