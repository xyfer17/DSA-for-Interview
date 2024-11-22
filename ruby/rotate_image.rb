# @param {Integer[][]} matrix
# @return {Void} Do not return anything, modify matrix in-place instead.
def rotate(matrix)

    for j in 0...matrix.length

        s = 0
        e = matrix.length - 1

        while s<e do

            temp = matrix[s][j]
            matrix[s][j] = matrix[e][j]
            matrix[e][j] = temp

            s+=1
            e-=1

        end

    end

    i = 0

    while i<matrix.length do
        j =0

        while j<i do
        
            temp = matrix[i][j]
            matrix[i][j] = matrix[j][i]
            matrix[j][i] = temp

            j+=1

        end
        
        i+=1

    end
   
end