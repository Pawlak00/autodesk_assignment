From main directory:

1. To build image:
   
`docker build . -t <image_name>`

2. To run:
   
`docker run -v $(pwd)/<dir_with_input>:/examples -e FILE_PATH=<path_to_input_data> <image_name>`