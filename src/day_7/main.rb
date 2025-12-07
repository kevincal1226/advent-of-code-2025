class Day7
  def initialize(filename)
    @filename = filename.capitalize
    @grid = []
  end

  def read_input
    File.open(@filename, 'r') do |f|
      f.each_line do |line|
        @grid.append(line[0..line.length - 2])
      end
    end

    @num_rows = @grid.length
    @num_cols = @grid[0].length
    @memo = Array.new(@num_rows) { Array.new(@num_cols, -1) }
    @memo_forreal = Array.new(@num_rows) { Array.new(@num_cols, 0) }
    @discovered_forreal = Array.new(@num_rows) { Array.new(@num_cols, false) }
  end

  def dfs_dp(row, col)
    return 0 if row == @num_rows || col < 0 || col >= @num_cols

    return @memo[row][col] unless @memo[row][col] == -1

    split_count = if @grid[row][col] == '^'
                    1 + dfs_dp(row, col + 1) + dfs_dp(row, col - 1)
                  else
                    dfs_dp(row + 1, col)
                  end

    @memo[row][col] = split_count

    @memo[row][col]
  end

  def part2
    row = 0
    col = @grid[0].index('S')
    puts dfs_dp(row, col) + 1
  end

  def dfs_forreal(row, col)
    return if row == @num_rows || col < 0 || col >= @num_cols || @discovered_forreal[row][col]

    @discovered_forreal[row][col] = true

    if @grid[row][col] == '^'
      @memo_forreal[row][col] = 1
      dfs_forreal(row, col + 1)
      dfs_forreal(row, col - 1)
    else
      dfs_forreal(row + 1, col)
    end
  end

  def part1forreal
    row = 0
    col = @grid[0].index('S')
    dfs_forreal(row, col)
    num_split = 0
    @memo_forreal.each do |row|
      num_split += row.inject(0, :+)
    end
    puts num_split
  end
end

g = Day7.new('input.txt')
g.read_input
g.part1forreal
g.part2
