# data model for cms project

## posts
- _id
- user_id
- created_at
- updated_at
- text
- media [optional]
- tags

## user
- _id
- first_name
- last_name
- created_at

## comment
- _id
- user_id
- post_id
- text
- created_at