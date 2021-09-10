# IRPF

A simple Go CLI to estimate the net salary in Spain given a raw annual amount.

**THIS IS NOT A FINANCIAL TOOL, USE AT YOUR OWN PERIL**

Example:

```shell
$ go build
$ irpf -h            
Usage of irpf:
  -a float
        raw annual income
  -p int
        number of individual payments (default 12)

$ irpf -a 28000 -p 14
Raw income: 28000 €
Retention: 6565.5 €
Retention percentage: 23.448214285714286 %
Net income: 21434.5 €
Net income per payment ( 14 pays): 1531.0357142857142 €
```
