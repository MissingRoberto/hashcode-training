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
  * Sort rides based on longest distance
  * For each rides, sort vehicules based on our greedy score.
```
time go run main.go resolver.go files/*
Score files/a_example.in: 12
Score files/b_should_be_easy.in: 169677
Score files/c_no_hurry.in: 16750973
Score files/d_metropolis.in: 14166512
Score files/e_high_bonus.in: 11588945
Total Score: 42676119
go run main.go resolver.go files/*  91.72s user 1.91s system 98% cpu 1:34.59 total
```

* Greedy: 
  * Sort rides based on longest distance and possible bonus
  * For each rides, sort vehicules based on our greedy score.

```
time go run main.go resolver.go files/*
Score files/a_example.in: 12
Score files/b_should_be_easy.in: 169677
Score files/c_no_hurry.in: 16750973
Score files/d_metropolis.in: 14166512
Score files/e_high_bonus.in: 11588945
Total Score: 42676119
go run main.go resolver.go files/*  174.67s user 1.78s system 99% cpu 2:57.59 total
```

It didn't affect previous result.

* Greedy: 
    * sort rides based on longest distance and bonus if possible.
    * For each ride, sort vehiclues based on our greedy score.

It timed out

* Greedy:
    * sort rides based on longest distance & "probability of bonus"
    * For each ride, sort vehicles based on our greedy score

```
time go run main.go resolver.go files/*
Score files/a_example.in: 12
Score files/b_should_be_easy.in: 169677
Score files/c_no_hurry.in: 16750973
Score files/d_metropolis.in: 14166512
Score files/e_high_bonus.in: 15738418
Total Score: 46825592
go run main.go resolver.go files/*  79.70s user 0.55s system 100% cpu 1:20.24 total
```

---
After fixing a couple of issues:
* Discard a ride if trip is not possible.

```
time go run main.go resolver.go files/*
Score files/a_example.in: 12
Score files/b_should_be_easy.in: 169977
Score files/c_no_hurry.in: 16750973
Score files/d_metropolis.in: 14166858
Score files/e_high_bonus.in: 16501528
Total Score: 47589348
go run main.go resolver.go files/*  83.19s user 0.62s system 100% cpu 1:23.70 total
```