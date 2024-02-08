package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"sort"
)

type Answer struct {
	AnswerText string
	EndingA    int
	EndingB    int
	EndingC    int
	EndingD    int
	EndingE    int
	EndingF    int
	EndingG    int
	EndingH    int
}

type Question struct {
	Text    string
	Answers []Answer
}

type Ending struct {
	Name      string
	Value     int
	Heading   string
	Paragraph string
	Img       string
	Alt       string
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/kviz", kvizHandler)

	http.HandleFunc("/submit", submitHandler)

	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func kvizHandler(w http.ResponseWriter, r *http.Request) {
	questions := []Question{
		{
			Text: "O kolik dětí se staráte?",
			Answers: []Answer{
				{"nemám děti", 0, 0, 0, 0, 0, 0, 0, 0},
				{"1 - 2", 30, 0, 0, 0, 0, 0, 0, 0},
				{"3 - 4", 50, 0, 0, 0, 0, 0, 0, 0},
				{"víc jak 4", 100, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			Text: "Kolik hodin naspíte za noc?",
			Answers: []Answer{
				{"0 - 3", 0, 0, 0, 0, 0, 0, 0, 0},
				{"4 - 6", 0, 20, 0, 0, 0, 0, 0, 0},
				{"7 - 9", 0, 30, 0, 0, 0, 0, 0, 0},
				{"víc jak 9", 0, 50, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			Text: "Kolik hodin týdně strávíte sportováním?",
			Answers: []Answer{
				{"nesportuji", 0, 0, 0, 0, 0, 0, 0, 0},
				{"1 - 3", 0, 0, 20, 0, 0, 0, 0, 0},
				{"4 - 6", 0, 0, 35, 0, 0, 0, 0, 0},
				{"víc jak 6", 0, 0, 50, 0, 0, 0, 0, 0},
			},
		},
		{
			Text: "Kolik hodin týdně strávíte nad svými koníčky?",
			Answers: []Answer{
				{"nemám koníčky", 0, 0, 0, 0, 0, 0, 0, 0},
				{"1 - 4", 0, 0, 0, 20, 0, 0, 0, 0},
				{"5 - 8", 0, 0, 0, 35, 0, 0, 0, 0},
				{"víc jak 8", 0, 0, 0, 50, 0, 0, 0, 0},
			},
		},
		{
			Text: "Kolik hodin týdně strávíte na sociálních sítích, když nejste v práci/škole?",
			Answers: []Answer{
				{"žiji pod kamenem", 0, 0, 0, 0, 0, 0, 0, 0},
				{"1 - 2", 0, 0, 0, 0, 0, 0, 0, 0},
				{"3 - 6", 0, 0, 0, 0, 0, 0, 10, 0},
				{"7 - 10", 0, 0, 0, 0, 0, 0, 30, 0},
				{"víc jak 10", 0, 0, 0, 0, 0, 0, 40, 0},
			},
		},
		{
			Text: "Jaký je váš průměrný denní screentime na mobilu za minulý týden?",
			Answers: []Answer{
				{"0 - 2", 0, 0, 0, 0, 0, 0, 0, 0},
				{"3 - 5", 0, 0, 0, 0, 0, 0, 20, 0},
				{"6 - 7", 0, 0, 0, 0, 0, 0, 35, 0},
				{"8 (plný úvazek)", 0, 0, 0, 0, 0, 0, 45, 0},
				{"8+ (přesčas)", 0, 0, 0, 0, 0, 0, 60, 0},
			},
		},
		{
			Text: "Kolik hodin denně strávíte hraním videoher?",
			Answers: []Answer{
				{"nehraju", 0, 0, 0, 0, 0, 0, 0, 0},
				{"0 - 2", 0, 0, 0, 0, 0, 0, 15, 0},
				{"3 - 4", 0, 0, 0, 0, 0, 0, 25, 0},
				{"4+", 0, 0, 0, 0, 0, 0, 35, 0},
			},
		},
		{
			Text: "Kolik hodin týdně strávíte v práci/škole, včetně studování doma a přesčasu?",
			Answers: []Answer{
				{"60 - 80 (otrok/hustler)", 0, 0, 0, 0, 0, 40, 0, 0},
				{"40 - 60 (pilný hoch)", 0, 0, 0, 0, 0, 30, 0, 0},
				{"30 - 40 (plný úvazek/student)", 0, 0, 0, 0, 0, 20, 0, 0},
				{"20 - 30 (polovičný úvazek)", 0, 0, 0, 0, 0, 10, 0, 0},
				{"mám bohaté rodiče", 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
		{
			Text: "Jste věřící?",
			Answers: []Answer{
				{"ano, ale nevěnuji tomu čas", 0, 0, 0, 0, 0, 0, 0, 0},
				{"ano, musím pravidelně chodit na schůzky", 0, 0, 0, 0, 0, 0, 0, 25},
				{"ne", 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}

	tmpl := template.Must(template.New("quiz").Funcs(template.FuncMap{"json": toJSON}).Parse(`
		<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta http-equiv="X-UA-Compatible" content="IE=edge">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Kvíz</title>
				<link href='https://fonts.googleapis.com/css?family=JetBrains Mono' rel='stylesheet'>
				<link rel="stylesheet" href="../static/styles.css"/>
				<script src="https://unpkg.com/htmx.org@1.9.6"></script>
			</head>
		<body>
			<h1><a href="/">nemascas.cz</a></h1>
				<div class="container">
					<h2>Dotazník:</h2>
					<div class="centered-element">
						
							<form action="/submit" method="post" id="myForm">
								{{range $index, $question := .}}
								<fieldset>
									<legend>{{$question.Text}}</legend>
									{{range $answerIndex, $answer := $question.Answers}}
										<label>
										<input type="radio" name="question_{{$index}}" value="{{. | json}}"> {{.AnswerText}}
										</label><br>
									{{end}}
								</fieldset>
								<br>
								{{end}}
					
								<button type="submit" value="Submit" hx-post="/submit" hx-target="#result">Jaké jsou zdroje mých časových nezmarů?</button>
								
							</form>
							
					</div>
					</div>

				<script>
					document.getElementById("myForm").addEventListener("submit", function(event) {
					event.preventDefault(); // Prevent the default form submission behavior
					// Your form submission handling logic here
					console.log("Form submitted");
					});
				</script>

				<div id="result" hx-swap="replace"></div>	
		</body>
		
		</html>	`))

	if err := tmpl.Execute(w, questions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	endings := []Ending{
		{"eA", 0, "Máš moc dětí", "prodej své děti", "static/images/A.jpg", "fotka dětí"},
		{"eB", 0, "Spíš moc hodin", "zkus spát 2 hodiny denně", "static/images/B.jpg", "fotka budíku"},
		{"eC", 0, "Sportuješ moc často", "zkus být tlustší", "static/images/C.jpg", "fotka sportovců"},
		{"eD", 0, "Máš moc koníčků", "zkus nemít život", "static/images/D.jpg", "fotka zábavných aktivit"},
		{"eE", 0, "", "", "", ""},
		{"eF", 0, "Pracuješ/studuješ až moc", "staň se bezdomovcem", "static/images/F.jpg", "fotka pracovníka"},
		{"eG", 0, "Jsi závislý na technologii", "přečti si manifesto Theodora Kaczynskiho", "static/images/G.jpg", "fotka technologie"},
		{"eH", 0, "Strávíš příliš času se svým stvořitelem", "nechoď do kostela, Bůh je i tak všude kolem tebe", "static/images/H.jpg", "fotka kostela"},
	}

	for key, values := range r.Form {
		if len(values) > 0 {
			if len(key) > len("question_") && key[:len("question_")] == "question_" {
				var selectedAnswer Answer
				err := json.Unmarshal([]byte(values[0]), &selectedAnswer)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				endings[0].Value += selectedAnswer.EndingA
				endings[1].Value += selectedAnswer.EndingB
				endings[2].Value += selectedAnswer.EndingC
				endings[3].Value += selectedAnswer.EndingD
				endings[4].Value += selectedAnswer.EndingE
				endings[5].Value += selectedAnswer.EndingF
				endings[6].Value += selectedAnswer.EndingG
				endings[7].Value += selectedAnswer.EndingH
			}
		}
	}

	if (endings[0].Value == 0) && (endings[1].Value == 0) && (endings[2].Value == 0) && (endings[3].Value == 0) && (endings[4].Value == 0) && (endings[5].Value == 0) && (endings[6].Value == 0) && (endings[7].Value == 0) {

		tmpl, err := template.ParseFiles("templates/prazdnyvysledky.html")

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	} else {

		var Filteredendings []Ending
		for _, v := range endings {
			if v.Value != 0 {
				Filteredendings = append(Filteredendings, v)
			}
		}

		sort.Slice(Filteredendings, func(i, j int) bool {
			return Filteredendings[i].Value > Filteredendings[j].Value
		})

		htmpl := `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Kvíz</title>
				<link rel="stylesheet" href="../static/styles.css"/>
			</head>
			<body>

			<div class="answercontainer">
				
				<ol>
					 {{range $index, $value := .}}
						  <li>
							<h3>{{indexPlusOne $index}}. {{ $value.Heading }}</h3>
							   <p>{{ $value.Paragraph }}</p>
							<img src="{{$value.Img}}" alt="{{$value.Alt}}">
						  </li>
					{{end}}
				</ol>
			
			</div>
			<div class="fakeend"></div>
			</body>
			</html>
			`

		tmpl, err := template.New("submit").Funcs(template.FuncMap{"indexPlusOne": indexPlusOne}).Parse(htmpl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, Filteredendings)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func toJSON(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func indexPlusOne(index int) int {
	return index + 1
}
