import re
import random

def get_theory_classroom_type(course_name, course_type):
    # 根据课程名称和类型选择合适的理论课教室
    special_rooms = {
        "音乐": "琴房",
        "舞蹈": "舞蹈教室",
        "美术": "美术教室",
        "声乐": "声乐教室",
    }
    
    for keyword, room in special_rooms.items():
        if keyword in course_name:
            return room
            
    # 随机选择普通教室或多媒体教室
    return random.choice(["普通教室", "多媒体教室"])

def get_experiment_classroom_type(course_name, course_type):
    # 根据课程名称和类型选择合适的实验课教室
    special_rooms = {
        "化工": "化工仿真实训室",
        "仿真": "虚拟仿真实训室",
        "自动化": "自动化生产线实训室",
        "液压": "液压气动实训室",
        "电工": "电工实训室",
        "网络": "网络实训室",
        "幼儿": "幼儿保健室",
        "机械": "机械装调实训室",
        "汽修": "汽修实训室",
    }
    
    for keyword, room in special_rooms.items():
        if keyword in course_name:
            return room
            
    return "实训车间"

def get_practice_classroom_type(course_name, course_type):
    # 根据课程名称和类型选择合适的实践课教室
    special_rooms = {
        "网络": "网络综合布线室",
        "幼儿": "蒙台梭立实训室",
        "活动": "活动室",
        "电气": "现代电气控制实训室",
    }
    
    for keyword, room in special_rooms.items():
        if keyword in course_name:
            return room
            
    return "实训车间"

def process_course_data(input_file, output_file):
    with open(input_file, 'r', encoding='utf-8') as f:
        content = f.read()

    # 找到所有的课程数据块
    pattern = r'is\.operate\.CourseLibraryData\(do\.CsCourseLibrary\{([^}]+)\}\)'
    course_blocks = list(re.finditer(pattern, content))
    
    for block in course_blocks:
        course_data = block.group(1)
        original_data = course_data
        
        # 提取课程信息
        name_match = re.search(r'Name:\s*"([^"]+)"', course_data)
        type_match = re.search(r'Type:\s*"([^"]+)"', course_data)
        course_name = name_match.group(1) if name_match else ""
        course_type = type_match.group(1) if type_match else ""
        
        # 提取各种学时
        theory_hours = re.search(r'TheoryHours:\s*(\d+)', course_data)
        experiment_hours = re.search(r'ExperimentHours:\s*(\d+)', course_data)
        practice_hours = re.search(r'PracticeHours:\s*(\d+)', course_data)
        computer_hours = re.search(r'ComputerHours:\s*(\d+)', course_data)
        credit_match = re.search(r'Credit:\s*([\d.]+)', course_data)
        credit = credit_match.group(1) if credit_match else "0"
        
        # 准备新的字段
        new_fields = []
        if theory_hours and int(theory_hours.group(1)) > 0:
            classroom = get_theory_classroom_type(course_name, course_type)
            new_fields.append(f'TheoryClassroomType: utils.Ptr("{classroom}"),')
        else:
            new_fields.append('TheoryClassroomType: nil,')
            
        if experiment_hours and int(experiment_hours.group(1)) > 0:
            classroom = get_experiment_classroom_type(course_name, course_type)
            new_fields.append(f'ExperimentClassroomType: utils.Ptr("{classroom}"),')
        else:
            new_fields.append('ExperimentClassroomType: nil,')
            
        if practice_hours and int(practice_hours.group(1)) > 0:
            classroom = get_practice_classroom_type(course_name, course_type)
            new_fields.append(f'PracticeClassroomType: utils.Ptr("{classroom}"),')
        else:
            new_fields.append('PracticeClassroomType: nil,')
            
        if computer_hours and int(computer_hours.group(1)) > 0:
            new_fields.append(f'ComputerClassroomType: utils.Ptr("机房"),')
        else:
            new_fields.append('ComputerClassroomType: nil,')
        
        # 删除已存在的教室类型字段
        course_data = re.sub(r'\s*TheoryClassroomType:\s*[^,]+,', '', course_data)
        course_data = re.sub(r'\s*ExperimentClassroomType:\s*[^,]+,', '', course_data)
        course_data = re.sub(r'\s*PracticeClassroomType:\s*[^,]+,', '', course_data)
        course_data = re.sub(r'\s*ComputerClassroomType:\s*[^,]+,', '', course_data)
        
        # 在 Credit 后添加新字段
        new_fields_str = '\n\t\t' + '\n\t\t'.join(new_fields)
        course_data = re.sub(r'Credit:\s*[\d.]+,?', f'Credit: {credit},', course_data)
        new_content = re.sub(r'Credit:\s*[\d.]+,?', f'Credit: {credit},{new_fields_str}', course_data)
        content = content.replace(original_data, new_content)
    
    with open(output_file, 'w', encoding='utf-8') as f:
        f.write(content)

# 处理文件
input_file = '/Users/xiaolfeng/ProgramProjects/Enterprise/FrontLeaves/ClassSchedulingSystem/table-install-cli/services/setup/setup_course.go'
output_file = '/Users/xiaolfeng/ProgramProjects/Enterprise/FrontLeaves/ClassSchedulingSystem/table-install-cli/services/setup/setup_course_new.go'
process_course_data(input_file, output_file) 