


upload()
{

	mod=$1
	proto_name=$2
	if [  "x$mod" = "x" ];then
		echo "your module-name is empty"
		exit 1
	fi

  set +e

	which npm
	if [  $? -ne  0 ];then
		echo -e "npm isn't installed, npm is installing..."
		yum install npm;
	fi

	which yapi-cli
	if [  $? -ne  0 ];then
		echo -e "yapi-cli isn't installed, npm is installing..."
		npm i -g yapi-cli --registry https://registry.npm.taobao.org
	fi

	set -e

	src=../../api/v1
	# if [ "x${token}" != "x" ];then
	# 	if [  ${#token}  -eq  64 ];then
	# 		sed -i "/^.*token/c\ \ \ \ \ \ \"token\": \"${token}\"," ${src}/yapi-import.json
	# 	else echo "token length does not match , automatically use the default token for you"
	# 	fi
	# fi

    echo $mod
    echo ${proto_name}
	#获取指定目录下的.proto的文件
  if [ "x${proto_name}" == "x" ];then
    for file in `ls -a $src`
      do
          if [ -f ${src}/$file ]
          then
        if [ "${file##*.}"x = "proto"x ] ;then
          name=${file%.*}.json
          cp ${src}/yapi-import.json  ${src}/${name}
          sed -i "/^.*file/c\ \ \ \ \ \ \"file\": \"$CWD/${src}/${file%.*}.swagger.json\"," ${src}/${name}
          pwd=$CWD/${src}/${name}
          yapi import --config=${pwd}
          fi
          fi
      done
  else
    name=${proto_name}.json
    cp ${src}/yapi-import.json  ${src}/${name}
    sed -i "/^.*file/c\ \ \ \ \ \ \"file\": \"$CWD/${src}/${proto_name}.swagger.json\"," ${src}/${name}
    pwd=${src}/${name}
    yapi import --config=${pwd}
  fi

}
upload $1 $2
