#!/bin/bash
cd db 
sql-migrate up
cd ..
./bin/cmd