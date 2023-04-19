# delete old proto 
rm -rf ./out/*

# compile proto files
for file in ./proto/*; do
    eval "$1 $file";
done