import os
import sys


limits = {
    "red": 12,
    "green": 13,
    "blue": 14,
}

def check_limits(cubes):
    for key, value in limits.items():
        if key not in cubes:
            continue
        if cubes[key] > value:
            return False

    return True

def check_game(line):
    # Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
    game_label, game_data = [x.strip() for x in line.split(":", 2)]
    game_id = int(game_label.split(" ")[-1])
    print(f"game_label = '{game_label}', game_data = '{game_data}', game_id: {game_id}")

    # 3 blue, 4 red
    games = [x.strip() for x in game_data.split(";")]

    for game in games:
        cubes = {}

        # 3 blue
        draws = [x.strip() for x in game.split(",")]

        for draw in draws:
            amount, color = draw.split(" ")
            amount = int(amount)
            cubes.setdefault(color, 0)
            cubes[color] += amount

        print(cubes)
        if not check_limits(cubes):
            return 0
    return game_id

def check_game_part2(line):
    # Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
    game_label, game_data = [x.strip() for x in line.split(":", 2)]
    game_id = int(game_label.split(" ")[-1])
    print(f"game_label = '{game_label}', game_data = '{game_data}', game_id: {game_id}")

    # 3 blue, 4 red
    games = [x.strip() for x in game_data.split(";")]

    cubes = {}

    for game in games:

        # 3 blue
        draws = [x.strip() for x in game.split(",")]

        for draw in draws:
            amount, color = draw.split(" ")
            amount = int(amount)
            if color not in cubes:
                cubes[color] = amount
            else:
                if amount > cubes[color]:
                    cubes[color] = amount

    result = 1
    for color, amount in cubes.items():
        result = result * amount

    print(result)
    return result

def main():
    if len(sys.argv) < 2:
        print("Need input file")
        return

    filename = sys.argv[1]
    result = 0

    with open(filename) as fd:
        for line in fd:
            line = line.strip()
            score = check_game_part2(line)
            result += score

    print(f"result = {result}")



if __name__ == "__main__":
    main()
