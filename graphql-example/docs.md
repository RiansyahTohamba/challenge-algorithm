# apa itu query?
ide graphql sama seperti SQL. 
bedanya pada syntax.

Instead of `SELECT`, GraphQL uses `Query` to request data.

# apa itu mutation?
syntax yang mengubah data pada SQL adalah  INSERT,UPDATE ,DELETE.
pada graphql disederhanakan menjadi `mutation`.

# apa itu subcription?
untuk listen data change melalui websocket.

# apa bedanya mutation vs query?
mutation untuk ubah vs query hanya read

# apa itu rootquery?
query{
    <!-- scope rootquery -->
    author {
        name
        books{
            title

        }
    }
    
}
nah author masuk kedalam rootquery sementara books tidak

