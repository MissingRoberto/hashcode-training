# 2018 Qualification
Basic algorithm:
```
Score files/a_example.in: 12
Score files/b_should_be_easy.in: 169677
Score files/c_no_hurry.in: 16750973
Score files/d_metropolis.in: 14166512
Score files/e_high_bonus.in: 11588945
Total Score: 42676119

0.47s user 1.14s system 108% cpu 1.481 total
```

* Greedy: Try to not wait and not picking rides which are further way
```
time go run main.go resolver.go files/*
Score files/a_example.in: 12
Score files/b_should_be_easy.in: 169827
Score files/c_no_hurry.in: 16750973
Score files/d_metropolis.in: 14166536
Score files/e_high_bonus.in: 15397392
Total Score: 46484740
go run main.go resolver.go files/*  197.59s user 1.38s system 98% cpu 3:22.43 total
```

* Greedy: 
  * Sort Rides by closest to previous one
  * Try to not wait and not picking rides which are further way
```
time go run main.go resolver.go files/*
Score files/a_example.in: 12
Score files/b_should_be_easy.in: 169827
Score files/c_no_hurry.in: 16750973
Score files/d_metropolis.in: 14166536
Score files/e_high_bonus.in: 15398301
Total Score: 46485649
go run main.go resolver.go files/*  202.31s user 1.38s system 99% cpu 3:25.68 total
```

* Greedy: 
  * For each rides, sort vehicules based on our greedy score.