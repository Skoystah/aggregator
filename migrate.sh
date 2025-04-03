#pushd 
cd $PWD/sql/schema/
#goose postgres "postgres://postgres:spitfire@localhost:5432/gator" down
goose postgres "postgres://postgres:spitfire@localhost:5432/gator" up
#popd
