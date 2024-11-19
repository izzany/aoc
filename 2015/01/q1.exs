fi = File.read!("./input.txt")

defmodule Counter do
  use Agent

  def start_link(initial_value) do
    Agent.start_link(fn -> initial_value end, name: __MODULE__)
  end

  def value do
    Agent.get(__MODULE__, & &1)
  end

  def increment do
    Agent.update(__MODULE__, &(&1 + 1))
  end

  def decrement do
    Agent.update(__MODULE__, &(&1 - 1))
  end
end

Counter.start_link(0)

count = fn a ->
  case a do
    "(" ->
      Counter.increment()

    ")" ->
      Counter.decrement()

    _ ->
      :ok
  end
end

for n <- String.graphemes(fi), do: count.(n)

IO.puts("total: #{Counter.value()}")
