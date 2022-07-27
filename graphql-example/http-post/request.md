ini menggunakan todo schema, silahkan buat schema lagi.

1. Retrieve single data of todo

```json
{ "query": "{ todo(id:\"b\") { id text done } }"  }
```
mengambil data todo dengan id='b'

```go
todo2 := schema.Todo{ID: "b", Text: "This is the most important", Done: false}
```

2. Retrieve todo List 
```json
{ "query": "{ todoList { id text done } }" }
```

3. Create New todo
```json
{ "query": "mutation { createTodo(text:\"My New todo\") { id text done } }" }
```

4. Update todo:
```json
{ "query": "mutation { updateTodo(id:\"a\", done: true) { id text done } }" }
```


# ==== 
```json
{ "query": "{ bookList { title } }" }
```
