DataTool是mac下可执行文件
DataTool.exe是windows下可执行文件

使用方法如下
例如:
    windows下:
		cd到exe文件所在目录
		输入 DataTool.exe F:\yongyu\jieguo "Total VOI volume" "Object volume" "Percent object vol"  回车
		输出结果类似:
			{F:/yongyu/jieguo\120__rec_tra.ctan.csv 120}
			{F:/yongyu/jieguo\121__rec_tra.ctan.csv 121}
			{F:/yongyu/jieguo\122__rec.ctan.csv 122}
			{F:/yongyu/jieguo\122__rec_tra.ctan.csv 122}

			cost=[694.1671ms]
			find your file at [F:/yongyu//newFile.csv]
		cost为所用时间
		[F:/yongyu//newFile.csv]为输出文件路径
		
		"Total VOI volume" "Object volume" "Percent object vol"为所要统计的列名，可以增加，修改或减少
	mac下:
		用法类似，cd到DataTool所在目录
		执行  DataTool F:\yongyu\jieguo "Total VOI volume" "Object volume" "Percent object vol"
		其他一样