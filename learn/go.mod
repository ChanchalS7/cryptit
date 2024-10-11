module example.com/learn

go 1.23.1

require github.com/pborman/uuid v1.2.1

require github.com/google/uuid v1.0.0 // indirect

replace github.com/ChanchalS7/cryptit/encrypt => ../cryptit/encrypt

replace github.com/ChanchalS7/cryptit/decrypt => ../cryptit/decrypt
