# TODO: complete for go web server container image
FROM scratch
COPY zips.csv zips.csv
#Copy executable file
ADD zipsvr zipsvr
EXPOSE 443
ENTRYPOINT ["/zipsvr"]
