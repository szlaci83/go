package main
import ("encoding/csv"
		"flag"
    	"fmt"
		"os"
		"time")

func main() {
	csvFilename := flag.String("csv", "problems.csv", "csv file, format : problem,answer")
	timeLimit := flag.Int("limit", 3, "the time limit in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err!= nil {
		exit(fmt.Sprintf("Filed to open CSV file :%s\n", *csvFilename))
		}
		r:= csv.NewReader(file)
		lines, err := r.ReadAll()
		if err!=nil {
			exit("Failed to parse CSV file.")
		}
		problems := parseLines(lines)

		timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

		correct :=0 
		for i, p := range problems {
			fmt.Printf("Problem #%d: %s = \n", i +1, p.q)
			answerCh := make(chan string)
			go func() {
				var answer string
				fmt.Scanf("%s\n", &answer)
				answerCh <- answer 
			}()
			select {
			case <-timer.C:
				fmt.Printf("You scored %d out of %d\n", correct, len(problems) )
				return
			case answer :=  <-answerCh:
				if answer == p.a {
					correct++
					}
			}
		}
	}

	func parseLines(lines [][] string) []problem{
		res := make ([]problem, len(lines))
		for i, line := range lines {
			res[i] = problem {
				q: line[0],
				a: line[1],
			}
		}
		return res 
	}

	type problem struct{
		q string
		a string
	}

	func exit(msg string){
		fmt.Println(msg)
		os.Exit(1)
	}


