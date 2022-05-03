from typing import List


class Solution:
    def maxArea(self, height: List[int]) -> int:
        l, r = 0, len(height) - 1
        area = 0

        while l < r:
            d = r - l
            t_height = 0
            if height[l] < height[r]:
                t_height = height[l]
                l += 1
            else:
                t_height = height[r]
                r -= 1

            t_area = d * t_height
            if t_area > area:
                area = t_area

        return area


solution = Solution()
print(solution.maxArea(height=[1, 8, 6, 2, 5, 4, 8, 3, 7]))
