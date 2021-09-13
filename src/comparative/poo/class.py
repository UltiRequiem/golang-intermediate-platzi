class PythonicEmployee:
    def __init__(self, id, name):
        self.id = id
        self.name = name

    def __str__(self):
        return f"{self.name} ({self.id}) is working with Python Code."


if __name__ == "__main__":
    me = PythonicEmployee(9999, "Zero")
    print(me)

    # Nice, also valid in js
    me.id = 45
    me.name = "Hi"

    print(me.id)
    print(me.name)
