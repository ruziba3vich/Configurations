In this project, I created a RESTful API for creating albums,
there are several endpoints have been implemented.
for example to get an album by the fields of `title`, `artist` or `price`, we should create a path
like `"http://127.0.0.1:5000/albums?price=$1"`, the endpoint has been shown in the `api` package.
Eventually, to get an `Album` bu title, if we need to use a freespace in the title, we should instead
use `%20`, because that's how we can represent a freespace in the path.
Finally, in the `GetAlbumByGivenIntervalHandler`, we should send the data in `JSON` format.
