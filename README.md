# GraphQL Server Example

This example shows how to implement a **GraphQL server with Golang** based on Prisma & [gqlgen](https://github.com/99designs/gqlgen).

## How to use

### 1. Download repo

Clone the repository:

```
git clone git@github.com:robojones/graphql.git
cd graphql/
```

### 2. Install the Prisma CLI

To run the example, you need the Prisma CLI. Please install it via Homebrew or [using another method](https://www.prisma.io/docs/prisma-cli-and-configuration/using-the-prisma-cli-alx4/#installation):

```
brew install prisma
brew tap
# or
npm i -g prisma
```

### 3. Set up database & deploy Prisma datamodel

Start the server and the database using docker-compose:

```bash
docker-compose up -d
```

Deploy our schema to our database:

```
prisma deploy # this also runs prisma generate and gqlgen
```

### 4. Start the GraphQL server

```
go run ./server
```

Navigate to [http://localhost:4000](http://localhost:4000) in your browser to explore the API of your GraphQL server in a [GraphQL Playground](https://github.com/prisma/graphql-playground).

### 5. Changing the GraphQL schema

After you made changes to prisma or graphql schema files, just generate the files:

```
go generate
```
