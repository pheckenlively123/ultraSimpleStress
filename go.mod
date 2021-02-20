module ultraSimpleStress

go 1.16

replace ultraSimpleStress/commandline => ./commandline
replace ultraSimpleStress/stresstests => ./stresstests

require ultraSimpleStress/commandline v1.2.3
require ultraSimpleStress/stresstests v1.2.3