# @param {Integer[]} prices
# @return {Integer}
def max_profit(prices)
    max_profit = 0

    buy_point = prices[0]

    for i in 1...prices.length

        val = prices[i] - buy_point

        if buy_point>prices[i]
            buy_point = prices[i]
        end

        if max_profit < val

            max_profit = val

        end

    end


    return max_profit

end
