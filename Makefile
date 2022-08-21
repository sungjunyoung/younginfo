all:
	docker build -t mypage ./mypage
	docker build -t ratings ./ratings
	docker build -t reviews ./reviews
	docker build -t details ./details