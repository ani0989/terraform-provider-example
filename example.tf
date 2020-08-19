provider "example" {
    username = "admin"
    password = "password"
}

resource "example_item" "test" {
    name = "Test"
}