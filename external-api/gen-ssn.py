from faker import Faker
fake = Faker()

with open("ssn.txt", "w") as text_file:
    for _ in range(1000):
        text_file.write("%s\n" % fake.ssn())
