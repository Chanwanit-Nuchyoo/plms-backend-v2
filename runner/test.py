def findMedianSortedArrays(nums1, nums2):
    if len(nums1) > len(nums2):
        nums1, nums2 = nums2, nums1

    x, y = len(nums1), len(nums2)

    low = 0
    high = x

    while low <= high:
        partitionX = (low + high) // 2
        partitionY = (x + y + 1) // 2 - partitionX

        maxLeftX = float("-inf") if partitionX == 0 else nums1[partitionX - 1]
        minRightX = float("inf") if partitionX == x else nums1[partitionX]

        maxLeftY = float("-inf") if partitionY == 0 else nums2[partitionY - 1]
        minRightY = float("inf") if partitionY == y else nums2[partitionY]

        if maxLeftX <= minRightY and maxLeftY <= minRightX:
            if (x + y) % 2 == 0:
                return float(max(maxLeftX, maxLeftY) + min(minRightX, minRightY)) / 2
            else:
                return max(maxLeftX, maxLeftY)

        elif maxLeftX > minRightY:
            high = partitionX - 1
        else:
            low = partitionX + 1


nums1 = list(
    map(int, input("Enter numbers for the first list, separated by space: ").split())
)
nums2 = list(
    map(int, input("Enter numbers for the second list, separated by space: ").split())
)
print(findMedianSortedArrays(nums1, nums2))
