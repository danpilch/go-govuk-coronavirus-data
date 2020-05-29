# GOVUK Coronavirus Deaths data

Formats death data in a table gathered from [GOVUK Coronavirus Data](https://c19downloads.azureedge.net/downloads/json/coronavirus-deaths_latest.json)

## Install

Download latest binary from [releases](https://github.com/danpilch/go-govuk-coronavirus-data/releases/)

# Usage
```
./go-govuk-coronavirus-data-linux-amd64 

+------------+--------+--------+
|    DATE    | DEATHS | CHANGE |
+------------+--------+--------+
| 2020-03-06 |      1 |      1 |
| 2020-03-07 |      2 |      1 |
| 2020-03-08 |      2 |      0 |
| 2020-03-09 |      3 |      1 |
...
| 2020-05-26 |  37048 |    134 |
| 2020-05-27 |  37460 |    412 |
| 2020-05-28 |  37837 |    377 |
| 2020-05-29 |  38161 |    324 |
+------------+--------+--------+
|              TOTAL  | 38161  |
+------------+--------+--------+
```
