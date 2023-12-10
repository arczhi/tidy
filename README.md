## TIDY 
### a efficient tool helps organizing your files !
#### feature
1. sort by time
2. sort by file type

#### usage
1. cmd

   ```
   tidy -h
   
   Usage of tidy:
     -dir string
           please give new a directory name (default "new_directory")
     -path string
           please give your directory path
     -time_span string
           give a time span, for example: 1d,1h,30min,30s
     -type string
           time: sort by time      file_type: sort by file_type (default "time")
   ```

2. code

   ```
   package main
   
   import (
   	"time"
   
   	tidy "github.com/arczhi/tidy/impl"
   	core "github.com/arczhi/tidy/pkg/core"
   )
   
   func main() {
   
   	// sort by time
   	t, err := tidy.New("./your/directory", core.WithTimeSpan(time.Duration(6)*time.Hour))
   	if err != nil {
   		panic(err)
   	}
   	if err := t.Exec(); err != nil {
   		panic(err)
   	}
   
   	//sort by file type
   	t2, err := tidy.New("./your/directory", core.WithFileType())
   	if err != nil {
   		panic(err)
   	}
   	if err := t2.Exec(); err != nil {
   		panic(err)
   	}
   }
   
   ```
   
   

#### example

![image-20231210225435370](README.assets/image-20231210225435370.png)

![image-20231210225520065](README.assets/image-20231210225520065.png)

![image-20231210225551409](README.assets/image-20231210225551409.png)
