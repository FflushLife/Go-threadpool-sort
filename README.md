### Intro
All commands below must be run from thread_pool_sort directory. <number_of_threads_and_cores> argument sets equaled number of goroutines and max(<number_of_threads_and_cores>, available_cores_count) number of cores.

### Installation

```sh
$ make go-install
```

### Preparing tests
```
python3 data_gen.py
```
This command creates file "rand_data.txt" with different numbers. This is one task. To create more tasks for thread pool run the script again. The tasks will be divided in the file by '\n' and simple numbers by ','. You can manually create this file with these rules and run the program.

### Run
```
$ ./bin/main <number_of_threads_and_cores>
```
All the sort results will be placed in "results" folder for checking if everything is correct.

### Conclusion
Conclusion may be watched from .ipynb file.
