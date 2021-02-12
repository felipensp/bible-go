# bible-go
A command-line interface for Holy Bible reading written in Go

```
bible.go book-name verse-reference [version]
```
##### Default version: acf

### Usage

#### Unique verse reference

```
> go run .\bible.go mt 7:21 acf
Mateus 7 (acf)
21. "Nem todo o que me diz: Senhor, Senhor! entrará no reino dos céus, mas aquele que faz a vontade de meu Pai, que está nos céus."
```

#### All chapter

```
> go run .\bible.go sl 150 acf 
Salmos 150 (acf)
1. "Louvai ao SENHOR. Louvai a Deus no seu santuário; louvai-o no firmamento do seu poder."
2. "Louvai-o pelos seus atos poderosos; louvai-o conforme a excelência da sua grandeza."
3. "Louvai-o com o som de trombeta; louvai-o com o saltério e a harpa."
4. "Louvai-o com o tamborim e a dança, louvai-o com instrumentos de cordas e com órgãos."
5. "Louvai-o com os címbalos sonoros; louvai-o com címbalos altissonantes."
6. "Tudo quanto tem fôlego louve ao Senhor. Louvai ao Senhor."
```
