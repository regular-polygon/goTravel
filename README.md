## Input
- CSV in the following format
- name, latitude, longitude
- provide coordinates in decimal degrees format
```
Santa Monica Pier,34.008384546300896,-118.49864296497657
San Diego Zoo,32.736333905970355,-117.15225973965737
Golden Gate Bridge,37.81972444641028,-122.47843105423168
```

## Usage
Build the program first.
```bash
go build -o tsp-solver ./cmd/tsp/main.go 
```
Then run it. 
If your input data is in a file named `places.csv` and it's in the same directory as the program, you can run the following command.
```bash
./tsp-solver ./places.csv
```

Use an `-e` flag if you want the solver to use Euclidean distance between the coordinates.
```bash
./tsp-solver ./places.csv -e
```