+++
title = "Go + gRPC + OPA - A Perfect Union - Part 3"
author = ["Yash Agarwal"]
date = 2019-02-18T06:47:15+05:30
images = []
tags = ["GRPC", "OPA"]
categories = ["Technical"]
draft = false
series = ["Go + gRPC + OPA"]
+++

I finished my last [post](/posts/2019/02/go-grpc-opa-a-perfect-union-part-2/) with the following issue -

> Now, here one problem arises, how to make sure that the search results will not return any book which the user is not
> authorized to access. We will solve this problem using OPA in the next and last post of this series.

Let's solve this issue now. We will use OPA's declarative language, Rego, to implement policies which will decide on the
basis of some user-provided data, which all objects are to be returned to the user.

We will also define a list of all the users who are part of this library. Here we are hardcoding this data, as I did
not want to waste my time in implementing a user registration service, but this functionality is not very important from
our point of view. We will require only one field from this users data - the `user_type` field. This field will
determine what the access level for the user is. We have already added the `access_level` field in the `Book` definition
of our proto file.

When the user wants to search for a particular book, it will provide its `user_type` the ISBN of the book to our service. Our service
will take that ISBN and pass it to the OPA server. OPA server already has the `Book` data and the `User` data. Now it has
the required ISBN to query the Book data. The Rego policy will query the Book data by ISBN. It will also
check for the `access_level` condition. Moreover, after this operation, it will return the resultant set of books that satisfy both the requirements.

Here is the Rego policy -

{{< highlight ruby >}}
package library

import data.books
import data.users
import input

search_books[book] {
  input.book.isbn == books[i].isbn
  input.user.user_type >= books[i].access_level
  book = books[i]
}

list_all_books[books[i]] {
  input.user.user_type >= books[i].access_level
}
{{< /highlight >}}

The user data is [here](https://github.com/yashhere/go-library-service/blob/master/OPA/users.json) and the book data is [here](https://github.com/yashhere/go-library-service/blob/master/add_books.sh).

A sample `input` request is shown below -

{{< highlight json >}}
{
    "input": {
        "book": {
            "isbn": "1128959038"
        },
        "user": {
            "user_type": 3
        }
    }
}

{{< /highlight >}}

The `input` is the data that the user is providing. In `search_books` function, the input ISBN is matched with the ISBN
of all books one by one. Then the resultant set of books is filtered by `user_type` and `access_level` (these
two fields are essentially the same). In the last, the resultant set of books is assigned to the variable `book` which
will be returned to the gRPC service.

The `list_all_books` function is implemented similarly. The only difference is that we do not need to filter the books
by ISBN. Filtering by `access_level` is enough.

Now our library service is completed. It is a very basic service. The intention was to show that the decision-making process can be offloaded to the OPA to reduce the complexity of the services. In this example, the advantages might not
be obvious, but in large production environments, where many services are running, it can make a significant
difference.

The code for this series can be found on my [Github](https://github.com/yashhere/go-library-service) account.

I hope you liked the article. Share your views and suggestions in the comments.

Thanks for reading. Cheers :)