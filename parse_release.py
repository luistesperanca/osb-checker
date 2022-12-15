import io
import sys


def main():
    release_notes = sys.argv[1]
    buffer = io.StringIO(release_notes)
    lines = buffer.readlines()

    bumps = []

    for line in lines[0:-3]:
        if line.startswith('* Bump'):
            bumps.append(line)
        else:
            print(line, end='')

    if len(bumps) > 0:
        print("\n\r## Dependencies")
        for bump in bumps:
            print(bump, end='')

    print("".join(lines[-3:]), end='')


if __name__ == '__main__':
    main()