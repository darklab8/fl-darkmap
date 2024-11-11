# Galaxy coordinates

defined in multiuniverse.ini, can be float values.
this map is obviously using for coordinates beginning, top-left corner. and we have a simple X, Y coordinates pointing where location of a dot should be

```ini
[sector]
mapping = sector01  ; "SIRIUS SECTOR"
; system = id, pos      ; always must match universe.ini position for ease of future development
system = br01, 2, 11
system = br02, 4, 10
system = br03, 3, 12
system = br04, 2, 9
system = br05, 1, 12
system = br06, 0.5, 9.5
system = br07, 3, 9
system = br08, 1, 5
```

# System coordinates

and in this case we clearly the case of heliocentric coordinates. the center of the map is beginning of coordinates 

Deput Munitions
```ini
[Object]
nickname = Li01_10
ids_name = 196775
pos = -92652.9763513834, 6.00000000000005, -12126.5455936093
rotate = -5.78812127218152E-16, -82.6219880193629, 5.7362615759299E-16
```

Battleship Missouri
```ini
[Object]
nickname = Li01_03
ids_name = 196768
pos = 51113.7, 50, 7638.1
```

Cell grid by default is equal to 34k
```ini
[system]
nickname = br02
file = systems\br02\br02.ini
pos = 4, 10
msg_id_prefix = gcs_refer_system_br02
visit = 0
strid_name = 196615
ids_info = 66102
NavMapScale = 1.5
```

size of grid can be smaller if NavMapScale is higher than 1.

# Galaxy system regions

We will use flstat region by default with matchin to some category in settings file
plus we will make possible overrides
