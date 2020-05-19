## Day 2: Build a Web Server

Build a web app that fullfills the following two endpoints:

`localhost:8080/current_time` returns the current Unix time.
```json
{
    current_time: 123456789
}
```

`/locahost:8080/pairs` returns a list of the pairs who wrote the code with some attributes.
```json
[
    {
        name: "Suhaib Malik",
        favorite_dessert: "Almond Bearclaw",
        likes_yellow: false,
    },
    {
        ...
    }
]
```
List at least two persons.
