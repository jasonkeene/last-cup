A way of tracking how many days have passed since you last consumed caffeine.

## Installation

```bash
$ go get github.com/jasonkeene/last-cup
```

Make sure `$GOPATH/bin` is on your `PATH`.

## Usage

Set the environment variable `LAST_CUP` to the date that you had your last
caffeinated beverage.

```bash
$ export LAST_CUP=2015-5-26
```

Put this somewhere in your .bashrc for persistence.

```bash
$ last-cup # first day
You've just started a long journey of caffeine abstinence.
Look forward to not having those fun withdrawl headaches I'm sure you're experiencing.

$ last-cup # after a while
Great job! It's been 3 months 2 days since your last cup! You're a Zen master!
```
