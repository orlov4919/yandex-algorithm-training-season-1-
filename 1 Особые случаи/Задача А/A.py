tRoom, tCond = map(int, input().split())
state = input()

if state == "fan" or tRoom >= tCond and state == "heat" or tRoom <= tCond and state == "freeze":
    print(tRoom)
else:
    print(tCond)