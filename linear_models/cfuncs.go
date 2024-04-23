//go:build go1.2
// +build go1.2

package linear_models

/*222

void libLinearPrintFunc(char *);

void golearn_liblinear_print_func_cgo(char *c) {
    libLinearPrintFunc(c);
}

*/
import "C"
