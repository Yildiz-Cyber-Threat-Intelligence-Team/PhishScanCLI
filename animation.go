package main

import (
    "fmt"
    "time"
    "os"
    "os/exec"
    "runtime"
)

func clearScreen() {
    switch runtime.GOOS {
    case "windows":
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
    default:
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func animateText(text string, delay time.Duration) {
    for _, char := range text {
        fmt.Print(string(char))
        time.Sleep(delay)
    }
    fmt.Println()
}

func main() {

    clearScreen()

    titles := []string{
        "Yıldız Siber Tehdit İstihbaratı Takımı - Siber VATAN",
    }

   
    for _, title := range titles {
        animateText(title, 100*time.Millisecond)
        time.Sleep(500 * time.Millisecond)
        clearScreen()
    }

    fmt.Println("Yıldız Siber Tehdit İstihbaratı Takımı - Siber VATAN")
    fmt.Println()
	
    fmt.Println(`@@@@@@@@@@@#+-:..+@@@@@@@@@@@@@@@@
@@@@@@@+.     =@@@@@@@@@@@@@@@@@@@
@@@@@:      .@@@@@@@@@#@@@%@@@@@@@
@@@%.      .@@@@@@@@@@=  =@@@@@@@@
@@:#       #@@@@@@@@%=-...:%@@@@@@
@:.%.     .@@@@@@@@@@@@=#@@@@@@@@@
*  ..      #@@@@@@@@@@@@@@@@@@@@@@
-   ..     .==---=*@@@@@@@@@@@*#@@
*                  ..*@@@@@@@*.*@@
-=              :+#@@*%@@@@#.  #@@
* .           .. :*%@@%%+.    .@@@
@-                            %@@@
@@:                         .#@@@@
@@@*.                      .@@@@@@
@@@@@=.           .--=--=+%@@@@@@@
@@@@@@@%:.  .:.       .=@@@@@@@@@@
@@@@@@@@@@@%#*+*#**%@@@@@@@@@@@@@@`)
    fmt.Println("Hadi balık tutalım...")
    fmt.Println()
}
