package yildizAnimation

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
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

func PrintAnimation() {

	clearScreen()

	titles := []string{
		"Yıldız Cyber Threat Intelligence Team - Siber VATAN",
	}

	for _, title := range titles {
		animateText(title, 100*time.Millisecond)
		time.Sleep(500 * time.Millisecond)
		clearScreen()
	}

	fmt.Println("Yıldız Cyber Threat Intelligence Team - Siber VATAN")
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
	fmt.Println("Let's go fishing...")
	fmt.Println()
}
