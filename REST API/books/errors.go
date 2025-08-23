package books

import "errors"

var ErrBookNotFound = errors.New("ошибка, книга, с таким названием несуществует!")
