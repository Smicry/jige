from agents.agents import get_agents


def main():
    while True:
        try:
            i = input("🤖 你好，我能为你做些什么？\n")
            print(f"收到请求：{i}")
            match i:
                case "quit" | "exit":
                    break
                case _:
                    get_agents()
        except KeyboardInterrupt:
            break


if __name__ == "__main__":
    main()
